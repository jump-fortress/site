<script lang="ts">
  import { formatRunTime } from '$lib/src/times';

  type Props = {
    run_time: number;
    tempus_time_id: number | undefined;
  };

  let { run_time, tempus_time_id }: Props = $props();
  let time = $derived(formatRunTime(run_time));
</script>

{#if tempus_time_id}
  <a
    class="flex h-full items-center hover:text-primary hover:underline"
    href="https://tempus2.xyz/records/{tempus_time_id}"
    target="_blank">{time}</a>
{:else}
  <span class="">{time}</span>
{/if}
