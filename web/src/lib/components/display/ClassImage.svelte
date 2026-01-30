<script lang="ts">
  import demo from '$lib/assets/tf/demo.png';
  import soldier from '$lib/assets/tf/soldier.png';

  type Props = {
    player_class: string;
  };

  let { player_class }: Props = $props();
</script>

{#if player_class === 'Soldier'}
  <img class="filter-lavender h-full object-cover" src={soldier} alt="" />
{:else}
  <img class="filter-lavender h-full object-cover" src={demo} alt="" />
{/if}
