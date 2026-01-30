<script lang="ts">
  type Props = {
    oerror: OpenAPIError;
  };

  let { oerror }: Props = $props();
</script>

{#if oerror}
  <div class="flex flex-col">
    <span class="text-error">{oerror.detail}</span>
    {#each oerror?.errors as error}
      <div class="flex gap-1">
        <span class="text-primary">{error.location}</span>
        <span>{error.message}</span>
      </div>
    {/each}
  </div>
{/if}
