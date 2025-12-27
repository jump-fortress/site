<script lang="ts">
  import RocketLauncher from '$lib/components/RocketLauncher.svelte';
  import type { PlayerClassPoints, PlayerPoints } from '$lib/schema';
  import { fade } from 'svelte/transition';

  type Props = {
    selected_class: string;
    points: PlayerPoints;
    launcher: string;
  };

  let { selected_class, points, launcher }: Props = $props();
</script>

<div class="relative flex gap-2 leading-6">
  <div class="flex flex-col bg-jfgray-900/75 px-2 py-0.5">
    <span class="text-sm opacity-75">last 3 monthlies</span>
    <div class="flex items-center gap-2 text-base">
      {#if selected_class === 'Soldier'}
        <span>{points.soldier.last_3_monthly}</span>
        <span class="ml-auto">(#1)</span>
      {:else}
        <span>{points.demo.last_3_monthly}</span>
        <span class="ml-auto">(#1)</span>
      {/if}
    </div>
  </div>
  <div class="flex flex-col bg-jfgray-900/75 px-2 py-0.5">
    <span class="text-sm opacity-75">last 9 MOTWs</span>
    <div class="flex items-center gap-2 text-base">
      {#if selected_class === 'Soldier'}
        <span>{points.soldier.last_9_motw}</span>
        <span class="ml-auto">(#1)</span>
      {:else}
        <span>{points.demo.last_9_motw}</span>
        <span class="ml-auto">(#1)</span>
      {/if}
    </div>
  </div>
  <div class="flex flex-col bg-jfgray-900/75 px-2 py-0.5">
    <span class="text-sm opacity-75">total points</span>
    <div class="flex items-center gap-2 text-base">
      {#if selected_class === 'Soldier'}
        <span>{points.soldier.total}</span>
        <span class="ml-auto">(#1)</span>
      {:else}
        <span>{points.demo.total}</span>
        <span class="ml-auto">(#1)</span>
      {/if}
    </div>
  </div>
  {#if selected_class === 'Soldier'}
    <div transition:fade class="absolute -top-2 -right-14 rotate-75 rotate-y-180">
      <RocketLauncher {launcher} />
    </div>
  {/if}
</div>
