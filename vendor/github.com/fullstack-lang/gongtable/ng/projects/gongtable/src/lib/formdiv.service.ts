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

import { FormDivDB } from './formdiv-db';

// insertion point for imports
import { FormEditAssocButtonDB } from './formeditassocbutton-db'
import { FormSortAssocButtonDB } from './formsortassocbutton-db'
import { FormGroupDB } from './formgroup-db'

@Injectable({
  providedIn: 'root'
})
export class FormDivService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  FormDivServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private formdivsUrl: string

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
    this.formdivsUrl = origin + '/api/github.com/fullstack-lang/gongtable/go/v1/formdivs';
  }

  /** GET formdivs from the server */
  getFormDivs(GONG__StackPath: string): Observable<FormDivDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<FormDivDB[]>(this.formdivsUrl, { params: params })
      .pipe(
        tap(),
		// tap(_ => this.log('fetched formdivs')),
        catchError(this.handleError<FormDivDB[]>('getFormDivs', []))
      );
  }

  /** GET formdiv by id. Will 404 if id not found */
  getFormDiv(id: number, GONG__StackPath: string): Observable<FormDivDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.formdivsUrl}/${id}`;
    return this.http.get<FormDivDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched formdiv id=${id}`)),
      catchError(this.handleError<FormDivDB>(`getFormDiv id=${id}`))
    );
  }

  /** POST: add a new formdiv to the server */
  postFormDiv(formdivdb: FormDivDB, GONG__StackPath: string): Observable<FormDivDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let FormFields = formdivdb.FormFields
    formdivdb.FormFields = []
    let CheckBoxs = formdivdb.CheckBoxs
    formdivdb.CheckBoxs = []
    let FormEditAssocButton = formdivdb.FormEditAssocButton
    formdivdb.FormEditAssocButton = new FormEditAssocButtonDB
    let FormSortAssocButton = formdivdb.FormSortAssocButton
    formdivdb.FormSortAssocButton = new FormSortAssocButtonDB
    let _FormGroup_FormDivs_reverse = formdivdb.FormGroup_FormDivs_reverse
    formdivdb.FormGroup_FormDivs_reverse = new FormGroupDB

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<FormDivDB>(this.formdivsUrl, formdivdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
	      formdivdb.FormFields = FormFields
	      formdivdb.CheckBoxs = CheckBoxs
        formdivdb.FormGroup_FormDivs_reverse = _FormGroup_FormDivs_reverse
        // this.log(`posted formdivdb id=${formdivdb.ID}`)
      }),
      catchError(this.handleError<FormDivDB>('postFormDiv'))
    );
  }

  /** DELETE: delete the formdivdb from the server */
  deleteFormDiv(formdivdb: FormDivDB | number, GONG__StackPath: string): Observable<FormDivDB> {
    const id = typeof formdivdb === 'number' ? formdivdb : formdivdb.ID;
    const url = `${this.formdivsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<FormDivDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted formdivdb id=${id}`)),
      catchError(this.handleError<FormDivDB>('deleteFormDiv'))
    );
  }

  /** PUT: update the formdivdb on the server */
  updateFormDiv(formdivdb: FormDivDB, GONG__StackPath: string): Observable<FormDivDB> {
    const id = typeof formdivdb === 'number' ? formdivdb : formdivdb.ID;
    const url = `${this.formdivsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let FormFields = formdivdb.FormFields
    formdivdb.FormFields = []
    let CheckBoxs = formdivdb.CheckBoxs
    formdivdb.CheckBoxs = []
    let FormEditAssocButton = formdivdb.FormEditAssocButton
    formdivdb.FormEditAssocButton = new FormEditAssocButtonDB
    let FormSortAssocButton = formdivdb.FormSortAssocButton
    formdivdb.FormSortAssocButton = new FormSortAssocButtonDB
    let _FormGroup_FormDivs_reverse = formdivdb.FormGroup_FormDivs_reverse
    formdivdb.FormGroup_FormDivs_reverse = new FormGroupDB

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<FormDivDB>(url, formdivdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
	      formdivdb.FormFields = FormFields
	      formdivdb.CheckBoxs = CheckBoxs
        formdivdb.FormGroup_FormDivs_reverse = _FormGroup_FormDivs_reverse
        // this.log(`updated formdivdb id=${formdivdb.ID}`)
      }),
      catchError(this.handleError<FormDivDB>('updateFormDiv'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in FormDivService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("FormDivService" + error); // log to console instead

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
