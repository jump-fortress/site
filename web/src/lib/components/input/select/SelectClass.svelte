<script lang="ts">
  import demo from '$lib/assets/tf/demo.png';
  import soldier from '$lib/assets/tf/soldier.png';

  type Props = {
    selected: string;
    onsubmit: (value: string) => void;
  };

  let { selected, onsubmit }: Props = $props();
</script>

<!-- bring your own container -->

<!-- soldier -->
<button
  class="cursor-pointer p-2 transition-opacity
  {selected === 'Demo' ? 'opacity-25 hover:opacity-75' : ''}"
  onclick={() => {
    onsubmit('Soldier');
    selected = 'Soldier';
  }}
  draggable="false">
  <img class="filter-lavender max-w-16 object-cover" src={soldier} alt="" />
</button>
<!-- demo -->
<button
  class="cursor-pointer p-2 transition-opacity
  {selected === 'Soldier' ? 'opacity-25 hover:opacity-75' : ''}"
  onclick={() => {
    onsubmit('Demo');
    selected = 'Demo';
  }}
  draggable="false">
  <img class="filter-lavender max-w-16 object-cover" src={demo} alt="" />
</button>
