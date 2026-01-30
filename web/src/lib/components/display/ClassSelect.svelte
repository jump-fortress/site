<script lang="ts">
  import demo from '$lib/assets/tf/demo.png';
  import soldier from '$lib/assets/tf/soldier.png';

  type Props = {
    player_class: string;
  };

  let { player_class = $bindable() }: Props = $props();
</script>

<!-- bring your own container -->
<!-- soldier -->
<button
  class="cursor-pointer p-2 transition-opacity
  {player_class === 'Demo' ? 'opacity-25 hover:opacity-75' : ''}"
  onclick={() => {
    player_class = 'Soldier';
  }}
  draggable="false">
  <img class="filter-lavender max-w-16 object-cover" src={soldier} alt="" />
</button>
<!-- demo -->
<button
  class="cursor-pointer p-2 transition-opacity
  {player_class === 'Soldier' ? 'opacity-25 hover:opacity-75' : ''}"
  onclick={() => {
    player_class = 'Demo';
  }}
  draggable="false">
  <img class="filter-lavender max-w-16 object-cover" src={demo} alt="" />
</button>
