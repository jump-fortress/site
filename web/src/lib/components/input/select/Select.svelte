<script lang="ts">
  import Response from '$lib/components/input/Response.svelte';
  import SubmitIcon from '$lib/components/input/SubmitIcon.svelte';
  import Option from '$lib/components/input/select/Option.svelte';
  import { fade } from 'svelte/transition';

  type Props = {
    type: 'text' | 'button';
    options: string[];
    placeholder?: string;
    message?: string;
    onsubmit: (value: string) => Promise<InputResponse>;
  };

  let { type, options, placeholder = '', message = '', onsubmit }: Props = $props();

  let value: string = $state('');
  let response: Promise<InputResponse> = $derived(
    Promise.resolve({ error: false, message: message })
  );

  let index: number = $state(0);

  // filtered for autofill
  let derivedOptions: string[] = $derived(options);
  let option: string = $derived(derivedOptions[index]);
  let showOptions: boolean = $state(false);

  let input: HTMLInputElement | undefined = $state();

  function clamp(n: number) {
    index = n < 0 ? 0 : n >= derivedOptions.length ? derivedOptions.length - 1 : n;
  }

  function focusInput() {
    if (input) {
      input.focus();
    }
  }

  // filter options for autofill
  $effect(() => {
    if (type === 'text') {
      derivedOptions = options.filter((option) => option.includes(value));
    }
  });
</script>

<input
  class="size-full px-2 text-left text-primary"
  {type}
  {placeholder}
  bind:value
  bind:this={input}
  onkeydown={(event) => {
    const key = event.key;

    if (showOptions) {
      if (key === 'Tab' || key === 'Enter') {
        event.preventDefault();
        value = option;
        showOptions = false;
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
        response = onsubmit(value);
        return;
      }

      showOptions = true;
    }
  }}
  onfocus={() => {
    showOptions = true;
  }}
  onclick={() => {
    showOptions = true;
  }}
  onfocusout={() => {
    showOptions = false;
  }} />

<!-- submit button & response -->
<button
  in:fade
  class="relative grid size-10 cursor-pointer place-content-center"
  label="submit"
  onclick={() => {
    if (value) {
      response = onsubmit(value);
    }
  }}>
  <SubmitIcon {response} />
</button>
<Response {response} />

<!-- options -->
{#if showOptions && derivedOptions.length}
  <div
    class="absolute top-full -left-px z-40 flex max-h-60 w-7/8 flex-col overflow-x-hidden overflow-y-auto rounded-b-box border border-content/50 bg-base-900 py-1 pr-3">
    {#each derivedOptions as option, i}
      <Option
        value={option}
        selected={i === index}
        onmousedown={(event: MouseEvent) => {
          index = i;
          value = option;
          showOptions = false;
          event.preventDefault();
          focusInput();
        }} />
    {/each}
  </div>
{/if}
