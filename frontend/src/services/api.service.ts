import { RequestsService } from './requests';
import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  public requests: RequestsService;

  constructor(private _http: HttpClient) {
    this.requests = new RequestsService(_http)
  }

}
