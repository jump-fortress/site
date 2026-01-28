<script lang="ts">
  type Props = {
    valid: Promise<boolean>;
  };

  let { valid }: Props = $props();
</script>

{#await valid}
  <span class="icon-[mdi--loading] animate-spin"></span>
{:then valid}
  <span class="icon-[mdi--send] {valid ? '' : 'text-error'}"></span>
{/await}
