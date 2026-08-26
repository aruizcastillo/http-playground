import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

// This store is the defailt template for now. The frontend currently
// calls the backend through the regular service layer; a typical Pinia store
// that uses that service will be added later as part of the HTTP API example.
export const useCounterStore = defineStore('counter', () => {
  const count = ref(0)
  const doubleCount = computed(() => count.value * 2)
  function increment() {
    count.value++
  }

  return { count, doubleCount, increment }
})
