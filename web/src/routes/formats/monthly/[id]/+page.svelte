<script lang="ts">
  import MonthlyHeader from '$lib/components/display/MonthlyHeader.svelte';

  let { data } = $props();
</script>

{#await data.monthly then monthly}
  {#if monthly}
    <MonthlyHeader {monthly} />
  {/if}
{/await}
