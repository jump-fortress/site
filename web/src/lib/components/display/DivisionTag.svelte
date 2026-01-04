<script lang="ts">
  interface Props {
    div: string | undefined;
  }

  let { div }: Props = $props();
  let lowercaseDiv = $derived(div?.toLowerCase());
</script>

{#if div}
  <span
    class="rounded-full border-division-{lowercaseDiv} border-x-2 text-division-{lowercaseDiv} size-fit bg-base-900/50 px-2 pb-0.5 text-sm">
    {div}
  </span>
{/if}
