<script lang="ts">
  interface Props {
    div: string | undefined;
    playerClass?: string;
  }

  let { div, playerClass = undefined }: Props = $props();
  let lowercaseDiv = $derived(div?.toLowerCase());
</script>

{#if div}
  <span
    class="rounded-box border-l-2 border-div-{lowercaseDiv} bg-base-900/90 px-2 text-center text-div-{lowercaseDiv} size-fit pb-px text-sm"
    >{div} {playerClass}</span>
{/if}
