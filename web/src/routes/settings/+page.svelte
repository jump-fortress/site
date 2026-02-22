<script lang="ts">
  import { Client } from '$lib/api/api.js';
  import Input from '$lib/components/input/Input.svelte';
  import SelectButtons from '$lib/components/input/SelectButtons.svelte';
  import Section from '$lib/components/layout/Section.svelte';
  import { ApiPaths, type Player } from '$lib/schema.js';

  import soldier from '$lib/assets/tf/soldier.png';
  import demo from '$lib/assets/tf/demo.png';
  import rl_mangler from '$lib/assets/tf/rl_mangler.png';
  import rl_original from '$lib/assets/tf/rl_original.png';
  import rl_stock from '$lib/assets/tf/rl_stock.png';
  import Button from '$lib/components/input/Button.svelte';
  import type { PageData } from './$types';
  import Select from '$lib/components/input/Select.svelte';
  import Errors from '$lib/components/input/Errors.svelte';
  import Content from '$lib/components/layout/Content.svelte';
  import { divs } from '$lib/helpers/divs';

  let { data }: { data: PageData } = $props();

  let oerror: OpenAPIError = $state(undefined);
</script>

{#if data.session}
  <Content>
    {#await Client.GET(ApiPaths.get_player_self)}
      <span></span>
    {:then { data: player }}
      {#if player}
        <Errors {oerror} />
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
              oerror = resp.error;
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
              oerror = resp.error;
              return resp.response.ok;
            }} />
          {#await Client.GET(ApiPaths.get_maps)}
            <span></span>
          {:then { data: maps }}
            {#if maps}
              <Select
                label="fav map"
                type="text"
                options={['none'].concat(maps.map((m) => m.name))}
                placeholder={player.map_pref}
                onsubmit={async (value) => {
                  let resp = await Client.GET(ApiPaths.update_map_pref, {
                    params: { path: { map_name: value } }
                  });
                  oerror = resp.error;
                  return resp.response.ok;
                }} />
            {/if}
          {/await}

          <Button
            onsubmit={async () => {
              let resp = await Client.POST(ApiPaths.update_avatar);
              oerror = resp.error;
              return resp.response.ok;
            }}>
            <span>update avatar</span>
            <span class="icon-[mdi--check]"></span>
          </Button>
        </Section>

        <Section label="requests">
          <Input
            label="alias change"
            type="text"
            placeholder={player.alias}
            onsubmit={async (value) => {
              let resp = await Client.POST(ApiPaths.submit_request, {
                params: { path: { request_kind: 'alias update', content: value } }
              });
              oerror = resp.error;
              return resp.response.ok;
            }} />

          <Select
            label="soldier div"
            type="text"
            placeholder={player.soldier_div}
            options={divs.concat('none')}
            onsubmit={async (value) => {
              let resp = await Client.POST(ApiPaths.submit_request, {
                params: { path: { request_kind: 'soldier div', content: value } }
              });
              oerror = resp.error;
              return resp.response.ok;
            }} />

          <Select
            label="demo div"
            type="text"
            placeholder={player.demo_div}
            options={divs.concat('none')}
            onsubmit={async (value) => {
              let resp = await Client.POST(ApiPaths.submit_request, {
                params: { path: { request_kind: 'demo div', content: value } }
              });
              oerror = resp.error;
              return resp.response.ok;
            }} />
        </Section>

        <Section label="connections">
          <Input
            max_width="max-w-80"
            label="Tempus ID"
            type="text"
            placeholder={player.tempus_id?.toString()}
            onsubmit={async (value) => {
              let resp = await Client.GET(ApiPaths.set_tempus_id, {
                params: { path: { tempus_id: parseInt(value) } }
              });
              oerror = resp.error;
              return resp.response.ok;
            }} />

          <Input
            label="Steam Trade URL"
            type="text"
            placeholder={player.trade_token}
            onsubmit={async (value) => {
              let resp = await Client.POST(ApiPaths.set_trade_token, {
                params: { path: { steam_trade_url: value } }
              });
              oerror = resp.error;
              return resp.response.ok;
            }} />
        </Section>
      {:else}
        <span>no player :(</span>
      {/if}
    {/await}
  </Content>
{/if}
