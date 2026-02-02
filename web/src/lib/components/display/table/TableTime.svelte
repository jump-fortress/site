<script lang="ts">
  import { formatDuration } from '$lib/helpers/times';
  import type { Time } from '$lib/schema';

  type Props = {
    time: Time;
  };

  let { time }: Props = $props();
  let runTime = $derived(formatDuration(time.duration));
</script>

{#if time.tempus_time_id}
  <a
    class="font-mono text-sm hover:text-primary hover:underline"
    href="https://tempus2.xyz/records/{time.tempus_time_id}"
    target="_blank">{runTime}</a>
{:else}
  <span class="font-mono text-sm">{runTime}</span>
{/if}
