<script lang="ts">
  type Props = {
    code: string | undefined;
    country?: string | undefined;
  };

  let { code, country = undefined }: Props = $props();
</script>

<!-- show nothing if there's no flag code -->
{#if code && country}
  <div class="flex items-center gap-1 text-sm text-content/75">
    {@render img(code)}
    <span>{country}</span>
  </div>
{:else if code}
  {@render img(code)}
{/if}

{#snippet img(code: string)}
  <img
    src="https://flagcdn.com/32x24/{code}.png"
    srcset="https://flagcdn.com/64x48/{code}.png 2x,
    https://flagcdn.com/96x72/{code}.png 3x"
    width="32"
    height="24"
    alt=""
    class="object-none" />
{/snippet}
