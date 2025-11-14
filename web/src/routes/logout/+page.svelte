<script lang="ts">
	import { goto, invalidateAll } from '$app/navigation';
	import { Client } from '$lib/internalApi';

	async function handleSignOut() {
		const { error } = await Client.POST('/internal/session/sign-out', {
			fetch: fetch
		});
		if (!error) {
			invalidateAll();
			goto('/');
		} else {
			console.error('error signing out: ', error);
		}
	}

	await handleSignOut();
</script>
