<script lang="ts">
  type Props = {
    value: string;
    selected: boolean;
    onmousedown: (event: MouseEvent) => void;
  };

  let { value, selected, onmousedown }: Props = $props();

  let button: HTMLButtonElement | undefined = $state(undefined);

  $effect(() => {
    if (selected && button) {
      button.scrollIntoView({ behavior: 'auto' });
    }
  });
</script>

<button bind:this={button} class="group flex w-full text-left" {onmousedown}>
  <span class="mx-1 w-full rounded-md pl-2 group-hover:bg-base-800 {selected ? 'bg-base-800' : ''}">
    {value}
  </span>
</button>
