<script lang="ts">
  import {
    datetimeToMs,
    formatDate,
    formatRelative,
    formatTime,
    specialMs
  } from '$lib/helpers/temporal';

  type Props = {
    datetime: string;
    player?: boolean;
  };

  let { datetime, player = false }: Props = $props();
  const ms = $derived(datetimeToMs(datetime));
</script>

<div class="relative">
  {#if player}
    <span class="peer"
      >{(specialMs.get('player') as number) > ms ? 'the beginning' : formatRelative(ms)}</span>
  {:else}
    <span class="peer">{formatRelative(ms)}</span>
  {/if}
  <div
    class="absolute -top-1 right-0 rounded-box bg-base-900 px-2 py-1 text-nowrap text-content opacity-0 drop-shadow-sm/100 drop-shadow-base-900 transition-opacity peer-hover:opacity-100 hover:opacity-100">
    <span>
      {formatDate(ms)}
      <span class="text-primary">{formatTime(ms)}</span>
    </span>
  </div>
</div>
