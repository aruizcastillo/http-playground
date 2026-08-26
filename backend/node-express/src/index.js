import express from 'express'
import cors from 'cors'
import 'dotenv/config'

const app = express()
const PORT = process.env.PORT || 3000
const FRONTEND_ORIGIN = process.env.FRONTEND_ORIGIN || 'http://localhost:5173'

app.use(cors({
  origin: FRONTEND_ORIGIN,
}))

app.use(express.json())

app.get('/api/message', (req, res) => {
  res.json({
    message: 'Hello from NODE-EXPRESS backend',
  })
})

app.post('/api/message', (req, res) => {
  const { message } = req.body

  res.json({
    received: message,
  })
})

app.listen(PORT, () => {
  console.log(`Server running on http://localhost:${PORT}`)
})