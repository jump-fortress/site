<script lang="ts">
  import { Client } from '$lib/api/api.js';
  import Input from '$lib/components/input/Input.svelte';
  import SelectButtons from '$lib/components/input/SelectButtons.svelte';
  import Section from '$lib/components/layout/Section.svelte';
  import { ApiPaths } from '$lib/schema.js';

  import soldier from '$lib/assets/tf/soldier.png';
  import demo from '$lib/assets/tf/demo.png';
  import rl_mangler from '$lib/assets/tf/rl_mangler.png';
  import rl_original from '$lib/assets/tf/rl_original.png';
  import rl_stock from '$lib/assets/tf/rl_stock.png';
  import Button from '$lib/components/input/Button.svelte';

  let { data } = $props();
  const player = $derived(data.player);

  let err: OpenAPIError = $state(undefined);
</script>

{#if player}
  {#if err}
    <div class="flex flex-col">
      <span class="text-error">{err.detail}</span>
      {#each err?.errors as error}
        <div class="flex gap-1">
          <span class="text-primary">{error.location}</span>
          <span>{error.message}</span>
        </div>
      {/each}
    </div>
  {/if}

  <Section label="profile">
    <SelectButtons
      label="fav class"
      options={[
        { src: soldier, value: 'Soldier' },
        { src: demo, value: 'Demo' }
      ]}
      selected={player.class_pref}
      onsubmit={async (value) => {
        let resp = await Client.GET(ApiPaths.update_class_pref, {
          params: { path: { player_class: value as 'Soldier' | 'Demo' } }
        });
        err = resp.error;
        return resp.response.ok;
      }} />

    <SelectButtons
      label="fav launcher"
      options={[
        { src: rl_stock, value: 'stock' },
        { src: rl_original, value: 'original' },
        { src: rl_mangler, value: 'mangler' }
      ]}
      withNone={true}
      selected={player.launcher_pref ?? 'none'}
      onsubmit={async (value) => {
        let resp = await Client.GET(ApiPaths.update_launcher_pref, {
          params: { path: { launcher: value as 'none' | 'stock' | 'original' | 'mangler' } }
        });
        err = resp.error;
        return resp.response.ok;
      }} />

    <span>under construction</span>
    <Input
      label="fav map"
      type="text"
      placeholder={player.map_pref}
      onsubmit={async (value) => {
        let resp = await Client.GET(ApiPaths.update_map_pref, {
          params: { path: { map_name: value } }
        });
        err = resp.error;
        return resp.response.ok;
      }} />

    <span>under construction</span>
    <Button
      onsubmit={async () => {
        let resp = await Client.GET(ApiPaths.readyz);
        err = resp.error;
        return resp.response.ok;
      }}>
      <span>update avatar</span>
      <span class="icon-[mdi--check]"></span>
    </Button>
  </Section>

  <Section label="requests">
    <span>under construction</span>
  </Section>

  <Section label="connections">
    <Input
      label="Tempus ID"
      type="text"
      placeholder={player.tempus_id?.toString()}
      onsubmit={async (value) => {
        let resp = await Client.GET(ApiPaths.set_tempus_id, {
          params: { path: { tempus_id: parseInt(value) } }
        });
        err = resp.error;
        return resp.response.ok;
      }} />

    <span>under construction</span>
    <Input
      label="Steam Trade URL"
      type="text"
      placeholder={player.trade_token}
      onsubmit={async (value) => {
        let resp = await Client.GET(ApiPaths.readyz);
        err = resp.error;
        return resp.response.ok;
      }} />
  </Section>
{/if}
