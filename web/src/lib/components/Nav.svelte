<script lang="ts">
	import { slide } from 'svelte/transition';
	import type { Session } from '$lib/schema';
	import steam_signin_button from '$lib/assets/static/sits_small.png';
	import { onMount } from 'svelte';
	import { invalidateAll } from '$app/navigation';

	type Props = {
		route: string;
		session?: Session | null;
	};

	let { route, session = null }: Props = $props();
	let showNavMenu = $state(false);

	// todo: is this still necessary for updating the navbar when logging in / out?
	onMount(() => {
		if (!session) {
			invalidateAll();
		}
	});
</script>

<svelte:document
	onmousedown={(e) => {
		// @ts-ignore
		if (!e.target.dataset.nav) {
			showNavMenu = false;
		}
	}}
/>

<nav
	class="bg-jfgray-900/75 fixed z-50 flex h-16 w-full select-none justify-center backdrop-blur-sm"
>
	<div class="w-5xl flex items-center justify-between transition-colors">
		<!-- left nav -->
		<div class="hover:text-ctp-lavender-50/75 flex h-full grow">
			{@render NavPage('')}
			{@render NavPage('ladder')}
			{@render NavPage('formats')}
			{@render NavPage('help')}
		</div>
		<!-- right nav -->
		<div class="relative flex h-16 flex-row-reverse items-center" data-nav="true">
			{#if !session}
				<a href="http://localhost:8000/internal/session/steam/discover">
					<img class="cursor-pointer" src={steam_signin_button} alt="" data-nav="true" />
				</a>
			{:else}
				<button
					class="group relative cursor-pointer px-4"
					onmousedown={() => {
						showNavMenu = true;
					}}
					data-nav="true"
				>
					<img
						class="group-hover:border-ctp-lavender/50 border-jfgray-800 size-14 rounded-full border-2 transition-colors"
						src={session.steamAvatarUrl}
						alt=""
						data-nav="true"
					/>
					{#if showNavMenu}
						<ul
							in:slide
							class="bg-jfgray-900/90 starting:border-b-ctp-lavender/0 border-ctp-lavender/50 absolute -top-2 right-0 -z-10 flex size-fit w-44 cursor-default flex-col gap-1 rounded-b-lg border-2 p-2 pt-16 text-start transition-colors delay-150"
							data-nav="true"
						>
							<a
								href="/players/{session.id}"
								data-nav="true"
								class="hover:bg-jfgray-700 span-ellipsis rounded-lg p-1 pl-2 transition-colors"
								>{session.displayName}</a
							>

							<hr class="border-jfgray-700 border-1" />

							{#if session.role === 'admin' || session.role === 'mod'}
								<a
									href="/moderation"
									data-nav="true"
									class="hover:bg-jfgray-700 span-ellipsis rounded-lg p-1 pl-2 transition-colors"
									>moderation</a
								>

								<hr class="border-jfgray-700 border-1" />
							{/if}

							<a
								href="/settings"
								data-nav="true"
								class="hover:bg-jfgray-700 rounded-lg p-1 pl-2 transition-colors">settings</a
							>
							<a
								href="/logout"
								data-sveltekit-preload-data="tap"
								data-nav="true"
								class="hover:bg-jfgray-700 rounded-lg p-1 pl-2 transition-colors"
							>
								logout
							</a>
						</ul>
					{/if}
				</button>
				<a
					class="text-ctp-lavender-50 flex items-center px-2 opacity-75 transition-opacity hover:opacity-100"
					href="/support"
					title="support"
				>
					<span class="icon-[ri--heart-line] size-6"></span>
				</a>
				<a
					class="text-ctp-lavender-50 flex items-center px-2 opacity-75 transition-opacity hover:opacity-100"
					href="https://discord.gg/tusBc64wnv"
					title="discord"
				>
					<span class="icon-[ri--discord-fill] size-6"></span>
				</a>
			{/if}
		</div>
	</div>
</nav>

<hr class="h-14 opacity-0" />

{#snippet NavPage(name: string)}
	<a
		href="/{name === '' ? '' : name}"
		class="hover:text-ctp-lavender-50 flex h-full items-center px-4 transition-colors"
		><div class="relative">
			<span>{name === '' ? 'home' : name}</span>
			{#if (route.includes(name) && name !== '') || (route === '' && name === '')}
				<hr
					class="text-ctp-lavender border-1 absolute left-0 right-0 m-auto flex w-11/12 rounded-full"
				/>
			{/if}
		</div></a
	>
{/snippet}
