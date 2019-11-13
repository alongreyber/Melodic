import Home from './routes/Home.svelte'
import About from './routes/About.svelte'
import SpotifyCallback from './routes/SpotifyCallback.svelte'

let routes = new Map()

// Exact path
routes.set('/', Home)
routes.set('/about', About)
routes.set('/spotify_callback', SpotifyCallback)

export default routes
