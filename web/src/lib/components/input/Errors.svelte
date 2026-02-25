<script lang="ts">
  import { fade } from 'svelte/transition';

  type Props = {
    oerror: OpenAPIError;
  };

  let { oerror }: Props = $props();
</script>

{#key oerror}
  {#if oerror}
    <div
      in:fade|global
      class="flex w-full max-w-xl flex-col border border-error/25 bg-error/25 p-1">
      <div class="flex gap-1">
        <span class="icon-[mdi--error-outline] text-error"></span>
        <span class="text-error">{oerror.status}</span>
        <span>{oerror.title}</span>
      </div>
      <span class="text-content">{oerror.detail}</span>
      {#each oerror?.errors as error}
        {#if error}
          <div class="text flex gap-1 text-content/75">
            <span>{error.location}</span>
            <span>{error.message}</span>
          </div>
        {/if}
      {/each}
    </div>
  {/if}
{/key}
