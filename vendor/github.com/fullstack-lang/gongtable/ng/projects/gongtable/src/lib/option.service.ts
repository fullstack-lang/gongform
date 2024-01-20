// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs'
import { Observable, of } from 'rxjs'
import { catchError, map, tap } from 'rxjs/operators'

import { OptionDB } from './option-db'
import { Option, CopyOptionToOptionDB } from './option'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class OptionService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  OptionServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private optionsUrl: string

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.optionsUrl = origin + '/api/github.com/fullstack-lang/gongtable/go/v1/options';
  }

  /** GET options from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<OptionDB[]> {
    return this.getOptions(GONG__StackPath, frontRepo)
  }
  getOptions(GONG__StackPath: string, frontRepo: FrontRepo): Observable<OptionDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<OptionDB[]>(this.optionsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<OptionDB[]>('getOptions', []))
      );
  }

  /** GET option by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<OptionDB> {
    return this.getOption(id, GONG__StackPath, frontRepo)
  }
  getOption(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<OptionDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.optionsUrl}/${id}`;
    return this.http.get<OptionDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched option id=${id}`)),
      catchError(this.handleError<OptionDB>(`getOption id=${id}`))
    );
  }

  /** POST: add a new option to the server */
  post(optiondb: OptionDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<OptionDB> {
    return this.postOption(optiondb, GONG__StackPath, frontRepo)
  }
  postOption(optiondb: OptionDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<OptionDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<OptionDB>(this.optionsUrl, optiondb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`posted optiondb id=${optiondb.ID}`)
      }),
      catchError(this.handleError<OptionDB>('postOption'))
    );
  }

  /** DELETE: delete the optiondb from the server */
  delete(optiondb: OptionDB | number, GONG__StackPath: string): Observable<OptionDB> {
    return this.deleteOption(optiondb, GONG__StackPath)
  }
  deleteOption(optiondb: OptionDB | number, GONG__StackPath: string): Observable<OptionDB> {
    const id = typeof optiondb === 'number' ? optiondb : optiondb.ID;
    const url = `${this.optionsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<OptionDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted optiondb id=${id}`)),
      catchError(this.handleError<OptionDB>('deleteOption'))
    );
  }

  // updateFront copy option to a version with encoded pointers and update to the back
  updateFront(option: Option, GONG__StackPath: string): Observable<OptionDB> {
    let optionDB = new OptionDB
    CopyOptionToOptionDB(option, optionDB)
    const id = typeof optionDB === 'number' ? optionDB : optionDB.ID
    const url = `${this.optionsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<OptionDB>(url, optionDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<OptionDB>('updateOption'))
    );
  }

  /** PUT: update the optiondb on the server */
  update(optiondb: OptionDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<OptionDB> {
    return this.updateOption(optiondb, GONG__StackPath, frontRepo)
  }
  updateOption(optiondb: OptionDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<OptionDB> {
    const id = typeof optiondb === 'number' ? optiondb : optiondb.ID;
    const url = `${this.optionsUrl}/${id}`;

    // insertion point for reset of pointers (to avoid circular JSON)
    // and encoding of pointers

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<OptionDB>(url, optiondb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`updated optiondb id=${optiondb.ID}`)
      }),
      catchError(this.handleError<OptionDB>('updateOption'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in OptionService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("OptionService" + error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {
    console.log(message)
  }
}
