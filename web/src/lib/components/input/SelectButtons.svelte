<script lang="ts">
  type Props = {
    label: string;
    options: {
      src: string;
      value: string;
    }[];
    withNone?: boolean;
    selected: string;
    onsubmit: (value: string) => Promise<boolean>;
  };

  let { label, options, withNone = false, selected, onsubmit }: Props = $props();

  let valid: Promise<boolean> = $state(Promise.resolve(true));
</script>

<!-- container for input -->
<label
  class="relative mt-1.5 flex w-fit rounded-box border border-base-700 bg-base-800 transition-colors hover:border-content/50">
  <!-- label -->
  <span class="absolute -top-1 left-2 bg-base-800 px-1 text-sm leading-1">{label}</span>
  <!-- input container -->
  <div class="flex h-18 w-full">
    {#each options as { src, value }}
      {@const isSelected: boolean = selected === value}
      <button
        class="cursor-pointer rounded-box
         {isSelected ? 'bg-base-900' : 'bg-base-800'}"
        onclick={async () => {
          const prev = selected;
          valid = onsubmit(value);
          selected = value;
          // rollback if invalid
          if (!(await valid)) {
            selected = prev;
          }
        }}>
        <img
          class="filter-lavender w-20 px-2 transition-opacity
          {isSelected ? '' : 'opacity-25 hover:opacity-75'}"
          {src}
          alt=""
          draggable="false" />
      </button>
    {/each}
    <!-- reset option -->
    {#if withNone}
      {@const isSelected: boolean = selected === 'none'}
      <!-- svelte-ignore a11y_consider_explicit_label -->
      <button
        class="grid w-20 cursor-pointer place-content-center rounded-box
        {isSelected ? 'bg-base-900' : 'bg-base-800'}"
        onclick={async () => {
          const prev = selected;
          valid = onsubmit('none');
          selected = 'none';
          // rollback if invalid
          if (!(await valid)) {
            selected = prev;
          }
        }}>
        <span class="icon-[mdi--close] size-6 {isSelected ? '' : 'opacity-50 hover:opacity-100'}"
        ></span>
      </button>
    {/if}
  </div>
</label>
