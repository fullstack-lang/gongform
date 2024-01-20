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

import { FormDivDB } from './formdiv-db'
import { FormDiv, CopyFormDivToFormDivDB } from './formdiv'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { FormFieldDB } from './formfield-db'
import { CheckBoxDB } from './checkbox-db'
import { FormEditAssocButtonDB } from './formeditassocbutton-db'
import { FormSortAssocButtonDB } from './formsortassocbutton-db'

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
    this.formdivsUrl = origin + '/api/github.com/fullstack-lang/gongform/go/v1/formdivs';
  }

  /** GET formdivs from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormDivDB[]> {
    return this.getFormDivs(GONG__StackPath, frontRepo)
  }
  getFormDivs(GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormDivDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<FormDivDB[]>(this.formdivsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<FormDivDB[]>('getFormDivs', []))
      );
  }

  /** GET formdiv by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormDivDB> {
    return this.getFormDiv(id, GONG__StackPath, frontRepo)
  }
  getFormDiv(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormDivDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.formdivsUrl}/${id}`;
    return this.http.get<FormDivDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched formdiv id=${id}`)),
      catchError(this.handleError<FormDivDB>(`getFormDiv id=${id}`))
    );
  }

  // postFront copy formdiv to a version with encoded pointers and post to the back
  postFront(formdiv: FormDiv, GONG__StackPath: string): Observable<FormDivDB> {
    let formdivDB = new FormDivDB
    CopyFormDivToFormDivDB(formdiv, formdivDB)
    const id = typeof formdivDB === 'number' ? formdivDB : formdivDB.ID
    const url = `${this.formdivsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<FormDivDB>(url, formdivDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<FormDivDB>('postFormDiv'))
    );
  }
  
  /** POST: add a new formdiv to the server */
  post(formdivdb: FormDivDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormDivDB> {
    return this.postFormDiv(formdivdb, GONG__StackPath, frontRepo)
  }
  postFormDiv(formdivdb: FormDivDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormDivDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    formdivdb.FormDivPointersEncoding.FormFields = []
    for (let _formfield of formdivdb.FormFields) {
      formdivdb.FormDivPointersEncoding.FormFields.push(_formfield.ID)
    }
    formdivdb.FormFields = []
    formdivdb.FormDivPointersEncoding.CheckBoxs = []
    for (let _checkbox of formdivdb.CheckBoxs) {
      formdivdb.FormDivPointersEncoding.CheckBoxs.push(_checkbox.ID)
    }
    formdivdb.CheckBoxs = []
    if (formdivdb.FormEditAssocButton != undefined) {
      formdivdb.FormDivPointersEncoding.FormEditAssocButtonID.Int64 = formdivdb.FormEditAssocButton.ID
      formdivdb.FormDivPointersEncoding.FormEditAssocButtonID.Valid = true
    }
    formdivdb.FormEditAssocButton = undefined
    if (formdivdb.FormSortAssocButton != undefined) {
      formdivdb.FormDivPointersEncoding.FormSortAssocButtonID.Int64 = formdivdb.FormSortAssocButton.ID
      formdivdb.FormDivPointersEncoding.FormSortAssocButtonID.Valid = true
    }
    formdivdb.FormSortAssocButton = undefined

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<FormDivDB>(this.formdivsUrl, formdivdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        formdivdb.FormFields = new Array<FormFieldDB>()
        for (let _id of formdivdb.FormDivPointersEncoding.FormFields) {
          let _formfield = frontRepo.FormFields.get(_id)
          if (_formfield != undefined) {
            formdivdb.FormFields.push(_formfield!)
          }
        }
        formdivdb.CheckBoxs = new Array<CheckBoxDB>()
        for (let _id of formdivdb.FormDivPointersEncoding.CheckBoxs) {
          let _checkbox = frontRepo.CheckBoxs.get(_id)
          if (_checkbox != undefined) {
            formdivdb.CheckBoxs.push(_checkbox!)
          }
        }
        formdivdb.FormEditAssocButton = frontRepo.FormEditAssocButtons.get(formdivdb.FormDivPointersEncoding.FormEditAssocButtonID.Int64)
        formdivdb.FormSortAssocButton = frontRepo.FormSortAssocButtons.get(formdivdb.FormDivPointersEncoding.FormSortAssocButtonID.Int64)
        // this.log(`posted formdivdb id=${formdivdb.ID}`)
      }),
      catchError(this.handleError<FormDivDB>('postFormDiv'))
    );
  }

  /** DELETE: delete the formdivdb from the server */
  delete(formdivdb: FormDivDB | number, GONG__StackPath: string): Observable<FormDivDB> {
    return this.deleteFormDiv(formdivdb, GONG__StackPath)
  }
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

  // updateFront copy formdiv to a version with encoded pointers and update to the back
  updateFront(formdiv: FormDiv, GONG__StackPath: string): Observable<FormDivDB> {
    let formdivDB = new FormDivDB
    CopyFormDivToFormDivDB(formdiv, formdivDB)
    const id = typeof formdivDB === 'number' ? formdivDB : formdivDB.ID
    const url = `${this.formdivsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<FormDivDB>(url, formdivDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<FormDivDB>('updateFormDiv'))
    );
  }

  /** PUT: update the formdivdb on the server */
  update(formdivdb: FormDivDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormDivDB> {
    return this.updateFormDiv(formdivdb, GONG__StackPath, frontRepo)
  }
  updateFormDiv(formdivdb: FormDivDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormDivDB> {
    const id = typeof formdivdb === 'number' ? formdivdb : formdivdb.ID;
    const url = `${this.formdivsUrl}/${id}`;

    // insertion point for reset of pointers (to avoid circular JSON)
    // and encoding of pointers
    formdivdb.FormDivPointersEncoding.FormFields = []
    for (let _formfield of formdivdb.FormFields) {
      formdivdb.FormDivPointersEncoding.FormFields.push(_formfield.ID)
    }
    formdivdb.FormFields = []
    formdivdb.FormDivPointersEncoding.CheckBoxs = []
    for (let _checkbox of formdivdb.CheckBoxs) {
      formdivdb.FormDivPointersEncoding.CheckBoxs.push(_checkbox.ID)
    }
    formdivdb.CheckBoxs = []
    if (formdivdb.FormEditAssocButton != undefined) {
      formdivdb.FormDivPointersEncoding.FormEditAssocButtonID.Int64 = formdivdb.FormEditAssocButton.ID
      formdivdb.FormDivPointersEncoding.FormEditAssocButtonID.Valid = true
    }
    formdivdb.FormEditAssocButton = undefined
    if (formdivdb.FormSortAssocButton != undefined) {
      formdivdb.FormDivPointersEncoding.FormSortAssocButtonID.Int64 = formdivdb.FormSortAssocButton.ID
      formdivdb.FormDivPointersEncoding.FormSortAssocButtonID.Valid = true
    }
    formdivdb.FormSortAssocButton = undefined

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<FormDivDB>(url, formdivdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        formdivdb.FormFields = new Array<FormFieldDB>()
        for (let _id of formdivdb.FormDivPointersEncoding.FormFields) {
          let _formfield = frontRepo.FormFields.get(_id)
          if (_formfield != undefined) {
            formdivdb.FormFields.push(_formfield!)
          }
        }
        formdivdb.CheckBoxs = new Array<CheckBoxDB>()
        for (let _id of formdivdb.FormDivPointersEncoding.CheckBoxs) {
          let _checkbox = frontRepo.CheckBoxs.get(_id)
          if (_checkbox != undefined) {
            formdivdb.CheckBoxs.push(_checkbox!)
          }
        }
        formdivdb.FormEditAssocButton = frontRepo.FormEditAssocButtons.get(formdivdb.FormDivPointersEncoding.FormEditAssocButtonID.Int64)
        formdivdb.FormSortAssocButton = frontRepo.FormSortAssocButtons.get(formdivdb.FormDivPointersEncoding.FormSortAssocButtonID.Int64)
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
