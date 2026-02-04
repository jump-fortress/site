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
        <span class="text-content">{oerror.detail}</span>
      </div>
      {#each oerror?.errors as error}
        <div class="text flex gap-1">
          <span class="text-content/75">{error.location}</span>
          <span>{error.message}</span>
        </div>
      {/each}
    </div>
  {/if}
{/key}
