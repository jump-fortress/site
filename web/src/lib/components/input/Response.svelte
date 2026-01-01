<script lang="ts">
  type Props = {
    response: Promise<InputResponse>;
  };

  let { response }: Props = $props();
</script>

{#await response then response}
  {#if response?.message}
    <span
      class="pointer-events-none absolute left-full ml-2 flex h-full items-center text-nowrap
          {response.error ? 'text-error' : 'text-primary'}">
      {response.message}
    </span>
  {/if}
{/await}
