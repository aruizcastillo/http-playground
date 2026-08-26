import axios from 'axios'

const API_URL = import.meta.env.VITE_API_URL

export const fetchMessage = () =>
  axios.get(`${API_URL}/api/message`)

export const sendMessage = (message) =>
  axios.post(`${API_URL}/api/message`, { message })