import type { Player, TimeWithPlayer } from '$lib/schema';

export function formatDuration(duration: number): string {
  const minutes = Math.floor(duration / 60)
    .toString()
    .padStart(2, '0');
  const seconds = (duration % 60).toFixed(3).padStart(6, '0');

  return `${minutes}:${seconds}`;
}

export function validDuration(duration: string): boolean {
  return /^((\d{0,2}):)?(\d{2}).(\d{1,3})$/g.test(duration);
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

export const twTableGradients: Map<string, string> = new Map([
  ['r1', 'bg-linear-0 from-div-gold/25 to-div-gold/25 text-div-gold'],
  ['r2', 'bg-linear-0 from-div-silver/25 to-div-silver/25 text-div-silver'],
  ['r3', 'bg-linear-0 from-div-bronze/25 to-div-bronze/25 text-div-bronze'],
  ['t1', 'bg-linear-to-r from-div-gold/25 from-0% to-div-gold/0 to-75%'],
  ['t2', 'bg-linear-to-r from-div-silver/25 from-0% to-div-silver/0 to-75%'],
  ['t3', 'bg-linear-to-r from-div-bronze/25 from-0% to-div-bronze/0 to-75%'],
  ['rdiamond', 'bg-linear-0 from-div-diamond/25 to-div-diamond/25 text-div-diamond'],
  ['rplatinum', 'bg-linear-0 from-div-platinum/25 to-div-platinum/25 text-div-platinum'],
  ['rgold', 'bg-linear-0 from-div-gold/25 to-div-gold/25 text-div-gold'],
  ['rsilver', 'bg-linear-0 from-div-silver/25 to-div-silver/25 text-div-silver'],
  ['rbronze', 'bg-linear-0 from-div-bronze/25 to-div-bronze/25 text-div-bronze'],
  ['rsteel', 'bg-linear-0 from-div-steel/25 to-div-steel/25 text-div-steel'],
  ['rwood', 'bg-linear-0 from-div-wood/25 to-div-wood/25 text-div-wood'],
  ['tdiamond', 'bg-linear-to-r from-div-diamond/25 from-0% to-div-diamond/0 to-25%'],
  ['tplatinum', 'bg-linear-to-r from-div-platinum/25 from-0% to-div-platinum/0 to-25%'],
  ['tgold', 'bg-linear-to-r from-div-gold/25 from-0% to-div-gold/0 to-25%'],
  ['tsilver', 'bg-linear-to-r from-div-silver/25 from-0% to-div-silver/0 to-25%'],
  ['tbronze', 'bg-linear-to-r from-div-bronze/25 from-0% to-div-bronze/0 to-25%'],
  ['tsteel', 'bg-linear-to-r from-div-steel/25 from-0% to-div-steel/0 to-25%'],
  ['twood', 'bg-linear-to-r from-div-wood/25 from-0% to-div-wood/0 to-25%']
]);
