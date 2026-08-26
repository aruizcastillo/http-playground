const API_URL = import.meta.env.VITE_API_URL

export const fetchMessage = () =>
    fetch(`${API_URL}/api/message`)

export const sendMessage = (message) =>
    fetch(`${API_URL}/api/message`, {
        method: 'POST',
        headers: {
        'Content-Type': 'application/json',
        },
        body: JSON.stringify({ message }),
    })