<script setup lang="ts">
import { defineComponent, ref } from 'vue'

const searchQuery = ref('')
const values = ['apple', 'banana', 'cherry', 'grape', 'orange', 'peach']
const showDropdown = ref(false)
const filteredResults = ref<string[]>([])

const search = () => {
  const query = searchQuery.value.toLowerCase()
  filteredResults.value = values.filter((value) => value.toLowerCase().includes(query))
  showDropdown.value = true
}

const selectResult = (result: string) => {
  searchQuery.value = result
  showDropdown.value = false
}
</script>

<template>
  <div>
    <input type="text" v-model="searchQuery" @input="search" />
    <ul v-if="showDropdown" class="dropdown">
      <li v-for="result in filteredResults" :key="result" @click="selectResult(result)">
        {{ result }}
      </li>
    </ul>
  </div>
</template>

<style scoped>
.dropdown {
  list-style-type: none;
  padding: 0;
  margin: 0;
  border: 1px solid #ccc;
}

.dropdown li {
  padding: 5px;
  cursor: pointer;
}

.dropdown li:hover {
  background-color: #f2f2f2;
}
</style>
