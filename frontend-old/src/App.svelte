<!-- Navigation links, using the "link" action -->
<!-- Also, use the "active" action to add the "active" CSS class when the URL matches -->
<nav class="navbar is-primary" role="navigation">
  <div class="navbar-brand">
      <div class="navbar-item">
	  <img src="logo-white.png" alt="Melodic Logo" height="40" >
	  <span class="logo-text">
	  Melodic
	  </span>
      </div>
  </div>
  <div class="navbar-start">
    <a class="navbar-item" href="/" use:link use:active>Home</a>
    <a class="navbar-item" href="/about" use:link use:active>About</a>
  </div>
  <div class="navbar-end">
      <a class="navbar-item" href="/" >Our Socials!</a>
  </div>
</nav>
<div class="container" id="main-container">
    <!-- Display notifications -->
    <Notifications></Notifications>
    <!-- Show the router -->
    <Router {routes} />
</div>


<script>
// Import the router component
import Router from 'svelte-spa-router'
// Import the "link" action and the methods to control history programmatically from the same module, as well as the location store
import { link, replace } from 'svelte-spa-router'
import active from 'svelte-spa-router/active'
// Import the list of routes
import routes from './routes.js'
import Notifications from './components/Notifications.svelte';
</script>

<style>
#main-container {
    margin-top: 15px;
}
.logo-text {
    margin-left: 5px;
    font-family: "Georgia", Serif;
    font-weight: bold;
    font-size: 1.5rem;
}
</style>
