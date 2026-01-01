<script lang="ts">
  type Props = {
    value: string;
    selected: boolean;
    onmousedown: (value: string) => void;
  };

  let { value, selected, onmousedown } = $props();

  let option: HTMLButtonElement | undefined = $state();

  $effect(() => {
    if (selected && option) {
      option.scrollIntoView({ behavior: 'auto' });
    }
  });
</script>

<button bind:this={option} class="group flex w-full text-left" {onmousedown}>
  <span class="mx-1 w-full rounded-md pl-2 group-hover:bg-base-800 {selected ? 'bg-base-800' : ''}">
    {value}
  </span>
</button>
