<script lang="ts">
  import Label from '$lib/components/input/Label.svelte';
  import SubmitIcon from '$lib/components/input/SubmitIcon.svelte';
  import Option from '$lib/components/input/Option.svelte';
  import { fade } from 'svelte/transition';

  type Props = {
    label: string;
    type: 'text' | 'button';
    max_width?: string;
    options: string[];
    placeholder?: string;
    value?: string;
    withSubmit?: boolean;
    onsubmit: (value: string) => Promise<boolean>;
  };

  let {
    label = '',
    type,
    max_width = 'max-w-80',
    options,
    placeholder = '',
    value = $bindable(),
    withSubmit = true,
    onsubmit
  }: Props = $props();

  let valid: Promise<boolean> = $state(Promise.resolve(true));

  let index: number = $state(0);

  // filtered for autofill
  let autofillOptions: string[] = $derived(options);
  let pendingOption: string | undefined = $derived(autofillOptions[index]);
  let focusOptions: boolean = $state(false);

  let input: HTMLInputElement | undefined = $state(undefined);

  function clamp(n: number) {
    index = n < 0 ? 0 : n >= autofillOptions.length ? autofillOptions.length - 1 : n;
  }

  // filter options for autofill
  $effect(() => {
    if (type === 'text' && value) {
      autofillOptions = options.filter((option) => option.includes(value!));
    }
  });
</script>

<Label {label} {max_width}>
  <input
    class="size-full px-2 text-primary"
    {type}
    {placeholder}
    bind:value
    bind:this={input}
    onkeydown={async (event) => {
      const key = event.key;

      if (focusOptions) {
        if (key === 'Tab' || key === 'Enter') {
          event.preventDefault();
          value = pendingOption;
          focusOptions = false;
        } else if (key === 'ArrowDown') {
          event.preventDefault();
          clamp(index + 1);
        } else if (key === 'ArrowUp') {
          event.preventDefault();
          clamp(index - 1);
        }
      } else {
        if (key === 'Enter' && value) {
          event.preventDefault();
          valid = onsubmit(value);
          if (await valid) {
            placeholder = value;
            value = '';
          }
          return;
        }

        focusOptions = true;
      }
    }}
    onfocus={() => {
      focusOptions = true;
    }}
    onclick={() => {
      focusOptions = true;
    }}
    onfocusout={() => {
      focusOptions = false;
    }} />

  <!-- submit button & response -->
  {#if withSubmit}
    <button
      in:fade
      class="relative grid size-9 cursor-pointer place-content-center"
      onclick={async () => {
        if (value) {
          valid = onsubmit(value);
          if (await valid) {
            placeholder = value;
            value = '';
          }
        }
      }}>
      <SubmitIcon {valid} />
    </button>
  {/if}

  <!-- options -->
  {#if focusOptions && autofillOptions.length}
    <div
      class="absolute top-full -left-px z-40 flex max-h-40 w-7/8 flex-col overflow-x-hidden overflow-y-auto rounded-b-box border border-t-0 border-content/50 bg-base-900 py-1 pr-3">
      {#each autofillOptions as option, i}
        <Option
          value={option}
          selected={i === index}
          onmousedown={(event: MouseEvent) => {
            index = i;
            value = option;
            focusOptions = false;
            event.preventDefault();
            if (input) {
              input.focus();
            }
          }} />
      {/each}
    </div>
  {/if}
</Label>
