export const divs: string[] = ['Diamond', 'Platinum', 'Gold', 'Silver', 'Bronze', 'Steel', 'Wood'];

// expects divs to have first letter uppercase
export function compareDivisions(a: string, b: string): number {
  let ai = divs.indexOf(a);
  let bi = divs.indexOf(b);
  if (ai === -1) {
    ai = 10;
  }
  if (bi === -1) {
    bi = 10;
  }
  if (ai === bi) return 0;
  return ai > bi ? 1 : -1;
}
