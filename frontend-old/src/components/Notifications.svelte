{#each myNotifications as n}
<div class="notification {n.color}">
    <button class="delete" on:click="{removeNotification(n)}"> </button>
    {n.text}
</div>
{/each}

<script>

let myNotifications;
import notificationStore from '../utils/notifications.js';
const unsubscribe = notificationStore.subscribe(list => myNotifications = list);

function removeNotification(nToRemove) {
    notificationStore.update(list => {
	return list.filter(n => n !== nToRemove);
    });
}
</script>
