<script lang="ts">
  import rl_mangler from '$lib/assets/tf/rl_mangler.png';
  import rl_original from '$lib/assets/tf/rl_original.png';
  import rl_stock from '$lib/assets/tf/rl_stock.png';

  const rlMap = new Map([
    ['stock', rl_stock],
    ['original', rl_original],
    ['mangler', rl_mangler]
  ]);

  type Props = {
    launcher: string | undefined;
  };

  let { launcher }: Props = $props();
</script>

{#if launcher}
  <img
    class="filter-lavender h-full object-cover"
    src={rlMap.get(launcher)}
    alt=""
    draggable="false" />
{/if}
