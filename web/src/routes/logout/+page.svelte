<script lang="ts">
  import { goto, invalidateAll } from '$app/navigation';
  import { Client } from '$lib/api/api';
  import { ApiPaths } from '$lib/schema';

  async function handleSignOut() {
    const { error } = await Client.POST(ApiPaths.sign_out, {
      fetch: fetch
    });
    if (!error) {
      invalidateAll();
      goto('/');
    } else {
      console.error('error signing out: ', error);
    }
  }

  handleSignOut();
</script>
