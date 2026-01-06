<script lang="ts">
  import demo from '$lib/assets/tf/demo.png';
  import soldier from '$lib/assets/tf/soldier.png';

  type Props = {
    selected: string;
  };

  let { selected } = $props();
</script>

{#if selected === 'Soldier'}
  <img class="filter-lavender h-full object-cover" src={soldier} alt="" />
{:else}
  <img class="filter-lavender h-full object-cover" src={demo} alt="" />
{/if}
