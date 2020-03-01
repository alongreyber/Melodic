import store from './store';
import router from './router';

export function getCookie(name) {
  var value = "; " + document.cookie;
  var parts = value.split("; " + name + "=");
  if (parts.length == 2) return parts.pop().split(";").shift();
}

var prefix = process.env.NODE_ENV == 'development' ? 'http://' : 'https://';

async function handleError(response) {
    let data = await response.json();
    if (!response.ok) {
	if(response.status == 401) {
	    store.commit('logOut');
	    router.push('/');
	    // Have to push first so this message doesn't get cleared
	    store.commit("displayMessage",
		{ color: 'is-warning', text: "Please log in" });
	} else {
	    // Some errors use the field "msg"
	    store.commit("displayMessage",
		{ color: 'is-danger', text: data.error });
	}
	return null;
    }
    return data.data;
}

export async function getApi(path) {
    const headers = new Headers();
    // Submit
    let result = await fetch(prefix + window.location.hostname + "/api" + path, {
	headers: headers,
	credentials: 'include'
    }).then(handleError);
    return result;
}

export async function postApi(path, body) {
    const headers = new Headers();
    headers.append("Content-Type", 'application/json');
    // Submit
    let result = await fetch(prefix + window.location.hostname + "/api" + path, {
	method: 'POST',
	headers: headers,
	credentials: 'include',
	body: JSON.stringify(body)
    }).then(handleError);
    return result;
}

export async function postApiForm(path, body) {
    const headers = new Headers();
    // Submit
    let result = await fetch(prefix + window.location.hostname + "/api" + path, {
	method: 'POST',
	headers: headers,
	credentials: 'include',
	body: body
    }).then(handleError);
    return result;
}

