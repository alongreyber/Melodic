import Home from './views/Home.svelte'
import About from './views/About.svelte'
import LoggedIn from './views/LoggedIn.svelte'

let routes = new Map()

// Exact path
routes.set('/', Home)
routes.set('/about', About)
routes.set('/logged_in', LoggedIn)

export default routes
