// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { CheckBoxDB } from './checkbox-db';

// insertion point for imports
import { FormDivDB } from './formdiv-db'

@Injectable({
  providedIn: 'root'
})
export class CheckBoxService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  CheckBoxServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private checkboxsUrl: string

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
    this.checkboxsUrl = origin + '/api/github.com/fullstack-lang/gongform/go/v1/checkboxs';
  }

  /** GET checkboxs from the server */
  getCheckBoxs(GONG__StackPath: string): Observable<CheckBoxDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<CheckBoxDB[]>(this.checkboxsUrl, { params: params })
      .pipe(
        tap(),
		// tap(_ => this.log('fetched checkboxs')),
        catchError(this.handleError<CheckBoxDB[]>('getCheckBoxs', []))
      );
  }

  /** GET checkbox by id. Will 404 if id not found */
  getCheckBox(id: number, GONG__StackPath: string): Observable<CheckBoxDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.checkboxsUrl}/${id}`;
    return this.http.get<CheckBoxDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched checkbox id=${id}`)),
      catchError(this.handleError<CheckBoxDB>(`getCheckBox id=${id}`))
    );
  }

  /** POST: add a new checkbox to the server */
  postCheckBox(checkboxdb: CheckBoxDB, GONG__StackPath: string): Observable<CheckBoxDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let _FormDiv_CheckBoxs_reverse = checkboxdb.FormDiv_CheckBoxs_reverse
    checkboxdb.FormDiv_CheckBoxs_reverse = new FormDivDB

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<CheckBoxDB>(this.checkboxsUrl, checkboxdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        checkboxdb.FormDiv_CheckBoxs_reverse = _FormDiv_CheckBoxs_reverse
        // this.log(`posted checkboxdb id=${checkboxdb.ID}`)
      }),
      catchError(this.handleError<CheckBoxDB>('postCheckBox'))
    );
  }

  /** DELETE: delete the checkboxdb from the server */
  deleteCheckBox(checkboxdb: CheckBoxDB | number, GONG__StackPath: string): Observable<CheckBoxDB> {
    const id = typeof checkboxdb === 'number' ? checkboxdb : checkboxdb.ID;
    const url = `${this.checkboxsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<CheckBoxDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted checkboxdb id=${id}`)),
      catchError(this.handleError<CheckBoxDB>('deleteCheckBox'))
    );
  }

  /** PUT: update the checkboxdb on the server */
  updateCheckBox(checkboxdb: CheckBoxDB, GONG__StackPath: string): Observable<CheckBoxDB> {
    const id = typeof checkboxdb === 'number' ? checkboxdb : checkboxdb.ID;
    const url = `${this.checkboxsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let _FormDiv_CheckBoxs_reverse = checkboxdb.FormDiv_CheckBoxs_reverse
    checkboxdb.FormDiv_CheckBoxs_reverse = new FormDivDB

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<CheckBoxDB>(url, checkboxdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        checkboxdb.FormDiv_CheckBoxs_reverse = _FormDiv_CheckBoxs_reverse
        // this.log(`updated checkboxdb id=${checkboxdb.ID}`)
      }),
      catchError(this.handleError<CheckBoxDB>('updateCheckBox'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in CheckBoxService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("CheckBoxService" + error); // log to console instead

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
