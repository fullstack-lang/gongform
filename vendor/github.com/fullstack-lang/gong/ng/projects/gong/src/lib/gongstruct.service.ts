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

import { GongStructDB } from './gongstruct-db'
import { GongStruct, CopyGongStructToGongStructDB } from './gongstruct'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { GongBasicFieldDB } from './gongbasicfield-db'
import { GongTimeFieldDB } from './gongtimefield-db'
import { PointerToGongStructFieldDB } from './pointertogongstructfield-db'
import { SliceOfPointerToGongStructFieldDB } from './sliceofpointertogongstructfield-db'

@Injectable({
  providedIn: 'root'
})
export class GongStructService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  GongStructServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private gongstructsUrl: string

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
    this.gongstructsUrl = origin + '/api/github.com/fullstack-lang/gong/go/v1/gongstructs';
  }

  /** GET gongstructs from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongStructDB[]> {
    return this.getGongStructs(GONG__StackPath, frontRepo)
  }
  getGongStructs(GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongStructDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<GongStructDB[]>(this.gongstructsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<GongStructDB[]>('getGongStructs', []))
      );
  }

  /** GET gongstruct by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongStructDB> {
    return this.getGongStruct(id, GONG__StackPath, frontRepo)
  }
  getGongStruct(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongStructDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.gongstructsUrl}/${id}`;
    return this.http.get<GongStructDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched gongstruct id=${id}`)),
      catchError(this.handleError<GongStructDB>(`getGongStruct id=${id}`))
    );
  }

  // postFront copy gongstruct to a version with encoded pointers and post to the back
  postFront(gongstruct: GongStruct, GONG__StackPath: string): Observable<GongStructDB> {
    let gongstructDB = new GongStructDB
    CopyGongStructToGongStructDB(gongstruct, gongstructDB)
    const id = typeof gongstructDB === 'number' ? gongstructDB : gongstructDB.ID
    const url = `${this.gongstructsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<GongStructDB>(url, gongstructDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<GongStructDB>('postGongStruct'))
    );
  }
  
  /** POST: add a new gongstruct to the server */
  post(gongstructdb: GongStructDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongStructDB> {
    return this.postGongStruct(gongstructdb, GONG__StackPath, frontRepo)
  }
  postGongStruct(gongstructdb: GongStructDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongStructDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    gongstructdb.GongStructPointersEncoding.GongBasicFields = []
    for (let _gongbasicfield of gongstructdb.GongBasicFields) {
      gongstructdb.GongStructPointersEncoding.GongBasicFields.push(_gongbasicfield.ID)
    }
    gongstructdb.GongBasicFields = []
    gongstructdb.GongStructPointersEncoding.GongTimeFields = []
    for (let _gongtimefield of gongstructdb.GongTimeFields) {
      gongstructdb.GongStructPointersEncoding.GongTimeFields.push(_gongtimefield.ID)
    }
    gongstructdb.GongTimeFields = []
    gongstructdb.GongStructPointersEncoding.PointerToGongStructFields = []
    for (let _pointertogongstructfield of gongstructdb.PointerToGongStructFields) {
      gongstructdb.GongStructPointersEncoding.PointerToGongStructFields.push(_pointertogongstructfield.ID)
    }
    gongstructdb.PointerToGongStructFields = []
    gongstructdb.GongStructPointersEncoding.SliceOfPointerToGongStructFields = []
    for (let _sliceofpointertogongstructfield of gongstructdb.SliceOfPointerToGongStructFields) {
      gongstructdb.GongStructPointersEncoding.SliceOfPointerToGongStructFields.push(_sliceofpointertogongstructfield.ID)
    }
    gongstructdb.SliceOfPointerToGongStructFields = []

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<GongStructDB>(this.gongstructsUrl, gongstructdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        gongstructdb.GongBasicFields = new Array<GongBasicFieldDB>()
        for (let _id of gongstructdb.GongStructPointersEncoding.GongBasicFields) {
          let _gongbasicfield = frontRepo.GongBasicFields.get(_id)
          if (_gongbasicfield != undefined) {
            gongstructdb.GongBasicFields.push(_gongbasicfield!)
          }
        }
        gongstructdb.GongTimeFields = new Array<GongTimeFieldDB>()
        for (let _id of gongstructdb.GongStructPointersEncoding.GongTimeFields) {
          let _gongtimefield = frontRepo.GongTimeFields.get(_id)
          if (_gongtimefield != undefined) {
            gongstructdb.GongTimeFields.push(_gongtimefield!)
          }
        }
        gongstructdb.PointerToGongStructFields = new Array<PointerToGongStructFieldDB>()
        for (let _id of gongstructdb.GongStructPointersEncoding.PointerToGongStructFields) {
          let _pointertogongstructfield = frontRepo.PointerToGongStructFields.get(_id)
          if (_pointertogongstructfield != undefined) {
            gongstructdb.PointerToGongStructFields.push(_pointertogongstructfield!)
          }
        }
        gongstructdb.SliceOfPointerToGongStructFields = new Array<SliceOfPointerToGongStructFieldDB>()
        for (let _id of gongstructdb.GongStructPointersEncoding.SliceOfPointerToGongStructFields) {
          let _sliceofpointertogongstructfield = frontRepo.SliceOfPointerToGongStructFields.get(_id)
          if (_sliceofpointertogongstructfield != undefined) {
            gongstructdb.SliceOfPointerToGongStructFields.push(_sliceofpointertogongstructfield!)
          }
        }
        // this.log(`posted gongstructdb id=${gongstructdb.ID}`)
      }),
      catchError(this.handleError<GongStructDB>('postGongStruct'))
    );
  }

  /** DELETE: delete the gongstructdb from the server */
  delete(gongstructdb: GongStructDB | number, GONG__StackPath: string): Observable<GongStructDB> {
    return this.deleteGongStruct(gongstructdb, GONG__StackPath)
  }
  deleteGongStruct(gongstructdb: GongStructDB | number, GONG__StackPath: string): Observable<GongStructDB> {
    const id = typeof gongstructdb === 'number' ? gongstructdb : gongstructdb.ID;
    const url = `${this.gongstructsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<GongStructDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted gongstructdb id=${id}`)),
      catchError(this.handleError<GongStructDB>('deleteGongStruct'))
    );
  }

  // updateFront copy gongstruct to a version with encoded pointers and update to the back
  updateFront(gongstruct: GongStruct, GONG__StackPath: string): Observable<GongStructDB> {
    let gongstructDB = new GongStructDB
    CopyGongStructToGongStructDB(gongstruct, gongstructDB)
    const id = typeof gongstructDB === 'number' ? gongstructDB : gongstructDB.ID
    const url = `${this.gongstructsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<GongStructDB>(url, gongstructDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<GongStructDB>('updateGongStruct'))
    );
  }

  /** PUT: update the gongstructdb on the server */
  update(gongstructdb: GongStructDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongStructDB> {
    return this.updateGongStruct(gongstructdb, GONG__StackPath, frontRepo)
  }
  updateGongStruct(gongstructdb: GongStructDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<GongStructDB> {
    const id = typeof gongstructdb === 'number' ? gongstructdb : gongstructdb.ID;
    const url = `${this.gongstructsUrl}/${id}`;

    // insertion point for reset of pointers (to avoid circular JSON)
    // and encoding of pointers
    gongstructdb.GongStructPointersEncoding.GongBasicFields = []
    for (let _gongbasicfield of gongstructdb.GongBasicFields) {
      gongstructdb.GongStructPointersEncoding.GongBasicFields.push(_gongbasicfield.ID)
    }
    gongstructdb.GongBasicFields = []
    gongstructdb.GongStructPointersEncoding.GongTimeFields = []
    for (let _gongtimefield of gongstructdb.GongTimeFields) {
      gongstructdb.GongStructPointersEncoding.GongTimeFields.push(_gongtimefield.ID)
    }
    gongstructdb.GongTimeFields = []
    gongstructdb.GongStructPointersEncoding.PointerToGongStructFields = []
    for (let _pointertogongstructfield of gongstructdb.PointerToGongStructFields) {
      gongstructdb.GongStructPointersEncoding.PointerToGongStructFields.push(_pointertogongstructfield.ID)
    }
    gongstructdb.PointerToGongStructFields = []
    gongstructdb.GongStructPointersEncoding.SliceOfPointerToGongStructFields = []
    for (let _sliceofpointertogongstructfield of gongstructdb.SliceOfPointerToGongStructFields) {
      gongstructdb.GongStructPointersEncoding.SliceOfPointerToGongStructFields.push(_sliceofpointertogongstructfield.ID)
    }
    gongstructdb.SliceOfPointerToGongStructFields = []

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<GongStructDB>(url, gongstructdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        gongstructdb.GongBasicFields = new Array<GongBasicFieldDB>()
        for (let _id of gongstructdb.GongStructPointersEncoding.GongBasicFields) {
          let _gongbasicfield = frontRepo.GongBasicFields.get(_id)
          if (_gongbasicfield != undefined) {
            gongstructdb.GongBasicFields.push(_gongbasicfield!)
          }
        }
        gongstructdb.GongTimeFields = new Array<GongTimeFieldDB>()
        for (let _id of gongstructdb.GongStructPointersEncoding.GongTimeFields) {
          let _gongtimefield = frontRepo.GongTimeFields.get(_id)
          if (_gongtimefield != undefined) {
            gongstructdb.GongTimeFields.push(_gongtimefield!)
          }
        }
        gongstructdb.PointerToGongStructFields = new Array<PointerToGongStructFieldDB>()
        for (let _id of gongstructdb.GongStructPointersEncoding.PointerToGongStructFields) {
          let _pointertogongstructfield = frontRepo.PointerToGongStructFields.get(_id)
          if (_pointertogongstructfield != undefined) {
            gongstructdb.PointerToGongStructFields.push(_pointertogongstructfield!)
          }
        }
        gongstructdb.SliceOfPointerToGongStructFields = new Array<SliceOfPointerToGongStructFieldDB>()
        for (let _id of gongstructdb.GongStructPointersEncoding.SliceOfPointerToGongStructFields) {
          let _sliceofpointertogongstructfield = frontRepo.SliceOfPointerToGongStructFields.get(_id)
          if (_sliceofpointertogongstructfield != undefined) {
            gongstructdb.SliceOfPointerToGongStructFields.push(_sliceofpointertogongstructfield!)
          }
        }
        // this.log(`updated gongstructdb id=${gongstructdb.ID}`)
      }),
      catchError(this.handleError<GongStructDB>('updateGongStruct'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in GongStructService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("GongStructService" + error); // log to console instead

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
