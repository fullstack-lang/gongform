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

import { FormGroupDB } from './formgroup-db'
import { FormGroup, CopyFormGroupToFormGroupDB } from './formgroup'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { FormDivDB } from './formdiv-db'

@Injectable({
  providedIn: 'root'
})
export class FormGroupService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  FormGroupServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private formgroupsUrl: string

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
    this.formgroupsUrl = origin + '/api/github.com/fullstack-lang/gongform/go/v1/formgroups';
  }

  /** GET formgroups from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormGroupDB[]> {
    return this.getFormGroups(GONG__StackPath, frontRepo)
  }
  getFormGroups(GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormGroupDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<FormGroupDB[]>(this.formgroupsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<FormGroupDB[]>('getFormGroups', []))
      );
  }

  /** GET formgroup by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormGroupDB> {
    return this.getFormGroup(id, GONG__StackPath, frontRepo)
  }
  getFormGroup(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormGroupDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.formgroupsUrl}/${id}`;
    return this.http.get<FormGroupDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched formgroup id=${id}`)),
      catchError(this.handleError<FormGroupDB>(`getFormGroup id=${id}`))
    );
  }

  // postFront copy formgroup to a version with encoded pointers and post to the back
  postFront(formgroup: FormGroup, GONG__StackPath: string): Observable<FormGroupDB> {
    let formgroupDB = new FormGroupDB
    CopyFormGroupToFormGroupDB(formgroup, formgroupDB)
    const id = typeof formgroupDB === 'number' ? formgroupDB : formgroupDB.ID
    const url = `${this.formgroupsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<FormGroupDB>(url, formgroupDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<FormGroupDB>('postFormGroup'))
    );
  }
  
  /** POST: add a new formgroup to the server */
  post(formgroupdb: FormGroupDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormGroupDB> {
    return this.postFormGroup(formgroupdb, GONG__StackPath, frontRepo)
  }
  postFormGroup(formgroupdb: FormGroupDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormGroupDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    formgroupdb.FormGroupPointersEncoding.FormDivs = []
    for (let _formdiv of formgroupdb.FormDivs) {
      formgroupdb.FormGroupPointersEncoding.FormDivs.push(_formdiv.ID)
    }
    formgroupdb.FormDivs = []

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<FormGroupDB>(this.formgroupsUrl, formgroupdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        formgroupdb.FormDivs = new Array<FormDivDB>()
        for (let _id of formgroupdb.FormGroupPointersEncoding.FormDivs) {
          let _formdiv = frontRepo.FormDivs.get(_id)
          if (_formdiv != undefined) {
            formgroupdb.FormDivs.push(_formdiv!)
          }
        }
        // this.log(`posted formgroupdb id=${formgroupdb.ID}`)
      }),
      catchError(this.handleError<FormGroupDB>('postFormGroup'))
    );
  }

  /** DELETE: delete the formgroupdb from the server */
  delete(formgroupdb: FormGroupDB | number, GONG__StackPath: string): Observable<FormGroupDB> {
    return this.deleteFormGroup(formgroupdb, GONG__StackPath)
  }
  deleteFormGroup(formgroupdb: FormGroupDB | number, GONG__StackPath: string): Observable<FormGroupDB> {
    const id = typeof formgroupdb === 'number' ? formgroupdb : formgroupdb.ID;
    const url = `${this.formgroupsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<FormGroupDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted formgroupdb id=${id}`)),
      catchError(this.handleError<FormGroupDB>('deleteFormGroup'))
    );
  }

  // updateFront copy formgroup to a version with encoded pointers and update to the back
  updateFront(formgroup: FormGroup, GONG__StackPath: string): Observable<FormGroupDB> {
    let formgroupDB = new FormGroupDB
    CopyFormGroupToFormGroupDB(formgroup, formgroupDB)
    const id = typeof formgroupDB === 'number' ? formgroupDB : formgroupDB.ID
    const url = `${this.formgroupsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<FormGroupDB>(url, formgroupDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<FormGroupDB>('updateFormGroup'))
    );
  }

  /** PUT: update the formgroupdb on the server */
  update(formgroupdb: FormGroupDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormGroupDB> {
    return this.updateFormGroup(formgroupdb, GONG__StackPath, frontRepo)
  }
  updateFormGroup(formgroupdb: FormGroupDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormGroupDB> {
    const id = typeof formgroupdb === 'number' ? formgroupdb : formgroupdb.ID;
    const url = `${this.formgroupsUrl}/${id}`;

    // insertion point for reset of pointers (to avoid circular JSON)
    // and encoding of pointers
    formgroupdb.FormGroupPointersEncoding.FormDivs = []
    for (let _formdiv of formgroupdb.FormDivs) {
      formgroupdb.FormGroupPointersEncoding.FormDivs.push(_formdiv.ID)
    }
    formgroupdb.FormDivs = []

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<FormGroupDB>(url, formgroupdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        formgroupdb.FormDivs = new Array<FormDivDB>()
        for (let _id of formgroupdb.FormGroupPointersEncoding.FormDivs) {
          let _formdiv = frontRepo.FormDivs.get(_id)
          if (_formdiv != undefined) {
            formgroupdb.FormDivs.push(_formdiv!)
          }
        }
        // this.log(`updated formgroupdb id=${formgroupdb.ID}`)
      }),
      catchError(this.handleError<FormGroupDB>('updateFormGroup'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in FormGroupService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("FormGroupService" + error); // log to console instead

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
