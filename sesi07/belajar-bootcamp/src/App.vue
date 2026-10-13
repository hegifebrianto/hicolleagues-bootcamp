<script setup>
import { ref } from 'vue';


import gambarVariable from './logo-hicolleagues.png'

const title = ref("HiColleagues Bootcamp on Variable Changes")
const gambar = "/src/logo-hicolleagues.png"

let count = ref(0)
function addCounter() {
  count.value++;
}
function minusCounter() {
  count.value--;
}
function resetCounter() {

  console.log(count.value)
  count.value = 0;
}
function resetName() {
  formName.value = ""
}

let activeClass = true;

let formName = ref("hegy")
</script>

<template>
  <h1>{{ title }}</h1>
  <h2>{{ activeClass }}</h2>
  <div>
    <!-- <button @click="addCounter">Add Counter</button>
    <button @click="minusCounter">Minus Counter</button>
    <button @click="resetCounter">Reset Counter</button> -->
    <!-- <button @click="resetName">Reset Name</button> -->
    <div v-bind:class="activeClass">
      <input type="text" name="formName" v-model="formName" id=""><br>

      <label for="">Siswa Hi Colleagues {{ formName }}</label><br>
      <img v-bind:src="gambar" width="200" height="200" class="image-profile" />
    </div>

    <!-- {{ count }} -->
  </div>



  <!-- <p>{{ gambarVariable }}</p>
  <p>{{gambar}}</p> -->
  <!-- <img src="./logo-hicolleagues.png" alt=""> -->
  <!-- <img v-bind:src="gambar" /> -->

  <!-- {{ count }} -->
  <!-- <di></div> -->
</template>

<style scoped>
.image-profile {
  /* margin-top: -5px; */
}
</style>
