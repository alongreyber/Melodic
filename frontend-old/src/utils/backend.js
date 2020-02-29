import notificationStore from '../utils/notifications.js';

export async function request(url, method = "GET") {
    const resp = await fetch(url, { credentials: "include"});
    const json = await resp.json();
    if(!resp.ok) {
	// Internal server error
	if(resp.status >= 500) {
	    notificationStore.update(list => {
		console.log("Writing error message");
		return list.append({
		    color: "is-danger",
		    text: "Encountered Server Error: " + json.error
		});
	    });
	    return;
	}
	else if(resp.status >= 400) {
	    notificationStore.update(list => {
		console.log("Writing error message");
		return list.append({
		    color: "is-warning",
		    text: "Encountered Client Error: " + json.error
		});
	    });
	    return;
	} else {
	    console.log("Unknown HTTP response");
	    return;
	}
    }
    return json.data;
}
