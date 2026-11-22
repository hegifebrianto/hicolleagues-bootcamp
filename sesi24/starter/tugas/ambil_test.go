package tugas

import (
	"context"
	"errors"
	"testing"
	"time"
)

// sumberLambat meniru pengambilan data yang butuh `lama`, menghormati ctx,
// dan gagal untuk id yang ada di `gagal`.
func sumberLambat(lama time.Duration, gagal map[int]error) Sumber {
	return func(ctx context.Context, id int) (Buku, error) {
		select {
		case <-time.After(lama):
		case <-ctx.Done():
			return Buku{}, ctx.Err()
		}
		if err := gagal[id]; err != nil {
			return Buku{}, err
		}
		return Buku{ID: id, Judul: "Buku " + string(rune('A'+id-1))}, nil
	}
}

func TestAmbilKonkuren_UrutanTerjaga(t *testing.T) {
	ids := []int{3, 1, 2, 5, 4}
	hasil, err := AmbilKonkuren(context.Background(), sumberLambat(10*time.Millisecond, nil), ids)
	if err != nil {
		t.Fatal(err)
	}
	for i, id := range ids {
		if hasil[i].ID != id {
			t.Errorf("indeks %d: ID %d, ingin %d (urutan harus sama dengan ids)", i, hasil[i].ID, id)
		}
	}
}

func TestAmbilKonkuren_LebihCepat(t *testing.T) {
	ids := []int{1, 2, 3, 4, 5}
	s := sumberLambat(60*time.Millisecond, nil)

	mulai := time.Now()
	if _, err := AmbilBerurutan(context.Background(), s, ids); err != nil {
		t.Fatal(err)
	}
	berurutan := time.Since(mulai) // ~300ms

	mulai = time.Now()
	if _, err := AmbilKonkuren(context.Background(), s, ids); err != nil {
		t.Fatal(err)
	}
	konkuren := time.Since(mulai) // ~60ms

	t.Logf("berurutan=%v konkuren=%v", berurutan, konkuren)
	if konkuren > berurutan/2 {
		t.Errorf("konkuren (%v) harusnya jauh lebih cepat dari berurutan (%v)", konkuren, berurutan)
	}
}

func TestAmbilKonkuren_ErrorMembatalkanLain(t *testing.T) {
	boom := errors.New("sumber rusak")
	s := sumberLambat(20*time.Millisecond, map[int]error{2: boom})

	_, err := AmbilKonkuren(context.Background(), s, []int{1, 2, 3})
	if !errors.Is(err, boom) {
		t.Fatalf("dapat %v, ingin %v", err, boom)
	}
}

func TestAmbilKonkuren_Timeout(t *testing.T) {
	ctx, batal := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer batal()

	mulai := time.Now()
	_, err := AmbilKonkuren(ctx, sumberLambat(2*time.Second, nil), []int{1, 2, 3})
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("dapat %v, ingin DeadlineExceeded", err)
	}
	if time.Since(mulai) > time.Second {
		t.Errorf("seharusnya berhenti segera saat timeout, tapi butuh %v", time.Since(mulai))
	}
}

func TestAmbilKonkuren_Kosong(t *testing.T) {
	hasil, err := AmbilKonkuren(context.Background(), sumberLambat(0, nil), nil)
	if err != nil || len(hasil) != 0 {
		t.Errorf("dapat %v, %v", hasil, err)
	}
}
