<script lang="ts">
  let { selected_class = $bindable() } = $props();

  import rocket from '$lib/assets/static/classes/rocket.png';
  import sticky from '$lib/assets/static/classes/sticky.png';
  import { Client } from '$lib/internalApi';
  import type { operations } from '$lib/schema';

  const style =
    'filter: brightness(0) saturate(100%) invert(97%) sepia(49%) saturate(6023%) hue-rotate(179deg) brightness(100%) contrast(108%)';

  async function handleUpdatePreferredClass(selectedClass: string) {
    const { response } = await Client.PUT('/internal/players/preferredclass/{class}', {
      fetch: fetch,
      params: {
        path: {
          class:
            selectedClass as operations['set-player-preferredclass']['parameters']['path']['class']
        }
      }
    });
  }
</script>

<div class="ml-auto flex h-full flex-col">
  <button
    onclick={() => {
      selected_class = 'Soldier';
      handleUpdatePreferredClass('Soldier');
    }}
    class="{selected_class === 'Soldier'
      ? 'opacity-100'
      : 'opacity-25 hover:opacity-75'} flex basis-1/2 cursor-pointer items-center p-3 pt-4 transition-all">
    <img class="size-16 select-none" {style} src={rocket} alt="" draggable="false" />
  </button>
  <button
    onclick={() => {
      selected_class = 'Demo';
      handleUpdatePreferredClass('Demo');
    }}
    class="{selected_class === 'Demo'
      ? 'opacity-100'
      : 'opacity-25 hover:opacity-75'} flex basis-1/2 cursor-pointer items-center p-3 pb-4 transition-all">
    <img class="size-16 select-none" {style} src={sticky} alt="" draggable="false" />
  </button>
</div>
