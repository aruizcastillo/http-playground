<script setup>
import { ref } from 'vue'
import {
  fetchMessage as fetchMessageWithFetch,
  sendMessage as sendMessageWithFetch,
} from './services/messageFetch.service'

import {
  fetchMessage as fetchMessageWithAxios,
  sendMessage as sendMessageWithAxios,
} from './services/messageAxios.service'

const getResponse = ref('')
const postMessage = ref('Hello from VUE frontend')
const postRequest = ref('')
const postResponse = ref('')

const handleFetchGet = async () => {
  const response = await fetchMessageWithFetch()
  const data = await response.json()

  getResponse.value = `FETCH: ${data.message}`
}

const handleAxiosGet = async () => {
  const response = await fetchMessageWithAxios()

  getResponse.value = `AXIOS: ${response.data.message}`
}

const handleFetchPost = async () => {
  const message = postMessage.value
  postRequest.value = message

  const response = await sendMessageWithFetch(message)
  const data = await response.json()

  postResponse.value = data.message
}

const handleAxiosPost = async () => {
  const message = postMessage.value
  postRequest.value = message

  const response = await sendMessageWithAxios(message)

  postResponse.value = response.data.message
}
</script>

<template>
  <main>
    <h1>HTTP Playground</h1>

    <section>
      <h2>GET /api/message</h2>

      <button @click="handleFetchGet">
        GET with fetch
      </button>

      <button @click="handleAxiosGet">
        GET with Axios
      </button>

      <p>Response: {{ getResponse }}</p>
    </section>

    <section>
      <h2>POST /api/message</h2>

      <input
        v-model="postMessage"
        type="text"
        placeholder="Write a message"
      />

      <button @click="handleFetchPost">
        POST with fetch
      </button>

      <button @click="handleAxiosPost">
        POST with Axios
      </button>

      <p>Request: {{ postRequest }}</p>
      <p>Response: {{ postResponse }}</p>
    </section>
  </main>
</template>
