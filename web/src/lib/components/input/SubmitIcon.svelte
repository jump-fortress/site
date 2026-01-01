<script lang="ts">
  type Props = {
    response: Promise<InputResponse>;
  };

  let { response }: Props = $props();
</script>

{#await response}
  <span class="icon-[mdi--loading] animate-spin"></span>
{:then response}
  {#if response}
    <span class="icon-[mdi--send] {response.error ? 'text-error' : ''}"></span>
  {:else}
    <span class="icon-[mdi--send]"></span>
  {/if}
{/await}
