import express from "express"
import morgan from "morgan"
import cors from "cors"
import recipeRouter from "./api/v1/router/recipe.router.js"

const app = express()

app.disable('x-powered-by')

app.use(morgan("dev"))
app.use(express.json())
app.use(express.urlencoded({ extended: true }))

app.use(
  cors({
    credentials: true,
    origin: 'http://localhost:5173'
  }),
)

app.use("/api/v1/recipe", recipeRouter)

export default app
