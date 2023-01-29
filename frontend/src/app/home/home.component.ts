import {Component, OnInit} from '@angular/core';
import {ApiService} from 'src/services/api.service';
import {Request} from "../../models/request";
import {formatDate} from "@angular/common";

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
})
export class HomeComponent implements OnInit {

  loading = false;
  requests: Request[] = [];
  displayColumns = ['createdAt', 'text', 'senderTgId']

  constructor(private api: ApiService) {
  }

  ngOnInit(): void {
    this.reloadData()

    setInterval(() => this.reloadData(), 2_000)
  }

  reloadData() {
    this.loading = true;
    this.api.requests.fetchAll(0, 250).subscribe(data => {
      const d = data.map(x => {
        x.senderTgId = `https://web.telegram.org/k/#${x.senderTgId}`
        x.createdAt = formatDate(x.createdAt, "dd.MM.yyyy HH:mm:ss", "en-US");
        return x
      }).filter((x) => {
        const ids = this.requests.map(y => y.id)
        return ids.find(y => y === x.id) !== null
      });

      this.requests = [...this.requests, ...d];

      this.loading = false;
    })
  }
}
