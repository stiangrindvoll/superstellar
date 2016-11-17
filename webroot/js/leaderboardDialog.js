export default class LeaderboardDialog {
  constructor () {
    this.domNode = document.getElementById('leaderboard');
  }

  show () {
    this.domNode.style.display = 'block'
  }

  hide() {
    this.domNode.style.display = 'none'
  }

  update(ranks) {
    let tbody = this.domNode.firstElementChild.firstElementChild;
    tbody.innerHTML = '';

    for(let rank of ranks) {
      let tr = this.buildRow(rank);
      tbody.appendChild(tr);
    }
  }

  buildRow(rank) {
    let tr = document.createElement("tr");

    tr.appendChild(this.buildCell(rank.rank, 'rank'));
    tr.appendChild(this.buildCell(rank.name, 'name'));
    tr.appendChild(this.buildCell(rank.score, 'score'));
    return tr;
  }

  buildCell(text, className) {
    let td = document.createElement("td");
    td.textContent = text;
    td.className = className;
    return td;
  }
}
