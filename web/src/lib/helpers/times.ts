import type { Player, TimeWithPlayer } from '$lib/schema';

export function formatDuration(duration: number) {
  const minutes = Math.floor(duration / 60)
    .toString()
    .padStart(2, '0');
  const seconds = (duration % 60).toFixed(3).padStart(6, '0');

  return `${minutes}:${seconds}`;
}

export function filterBestTimes(twps: TimeWithPlayer[]): TimeWithPlayer[] {
  // get unique players via set conversion
  const players: Player[] = Array.from(
    new Set(
      twps.map((twp) => {
        return twp.player;
      })
    )
  );

  // set each time to its best time
  const bestTimes: TimeWithPlayer[] = [];
  for (const p of players) {
    bestTimes.push(
      twps
        .filter(({ player }) => p.id === player.id)
        .reduce((prev, curr) => (prev.time.duration < curr.time.duration ? prev : curr))
    );
  }

  // remove duplicate times via another set conversion
  return Array.from(new Set(bestTimes));
}

export const twTimes: Map<string, string> = new Map([
  ['r1', 'bg-linear-0 from-div-gold/25 to-div-gold/25 text-div-gold'],
  ['r2', 'bg-linear-0 from-div-silver/25 to-div-silver/25 text-div-silver'],
  ['r3', 'bg-linear-0 from-div-bronze/25 to-div-bronze/25 text-div-bronze'],
  ['t1', 'bg-linear-to-r from-div-gold/25 from-0% to-div-gold/0 to-75%'],
  ['t2', 'bg-linear-to-r from-div-silver/25 from-0% to-div-silver/0 to-75%'],
  ['t3', 'bg-linear-to-r from-div-bronze/25 from-0% to-div-bronze/0 to-75%']
]);
