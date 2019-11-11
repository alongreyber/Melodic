import Home from './routes/Home.svelte'
import About from './routes/About.svelte'

let routes = new Map()

// Exact path
routes.set('/', Home)
routes.set('/about', About)

export default routes
