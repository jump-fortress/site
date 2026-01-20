<script lang="ts">
  import rl_mangler from '$lib/assets/tf/rl_mangler.png';
  import rl_original from '$lib/assets/tf/rl_original.png';
  import rl_stock from '$lib/assets/tf/rl_stock.png';

  const rlMap = new Map([
    ['Stock', rl_stock],
    ['Original', rl_original],
    ['Mangler', rl_mangler]
  ]);

  type Props = {
    launcher: string;
  };

  let { launcher }: Props = $props();
</script>

{#if launcher !== 'None'}
  <img
    class="filter-lavender h-full object-cover"
    src={rlMap.get(launcher)}
    alt=""
    draggable="false" />
{/if}
