<script lang="ts">
  interface Props {
    div: string | undefined;
  }

  let { div }: Props = $props();
  let lowercaseDiv = $derived(div?.toLowerCase());
</script>

{#if div}
  <span
    class="rounded-box border-l-2 border-division-{lowercaseDiv} bg-base-900/75 px-2 text-center text-division-{lowercaseDiv} size-fit pb-px text-sm"
    >{div}</span>
{/if}
